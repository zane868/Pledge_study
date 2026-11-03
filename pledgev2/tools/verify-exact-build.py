#!/usr/bin/env python3
"""Submit the exact full compiler input, avoiding via-IR single-source context differences."""
import json
import os
from pathlib import Path
import time
import urllib.parse
import urllib.request

root=Path(__file__).resolve().parent.parent
api_key=os.environ['ETHERSCAN_API_KEY']
endpoint='https://api.etherscan.io/v2/api?chainid=11155111'

def api(action, **fields):
    payload={'apikey':api_key,'module':'contract','action':action,**fields}
    req=urllib.request.Request(endpoint,data=urllib.parse.urlencode(payload).encode())
    with urllib.request.urlopen(req,timeout=90) as response:result=json.load(response)
    time.sleep(0.5)
    return result


def main():
    manifest_path=root/'networks/sepolia.json'
    manifest=json.loads(manifest_path.read_text())
    broadcast=json.loads((root/'broadcast/DeploySepolia.s.sol/11155111/run-latest.json').read_text())
    builds=[]
    for path in (root/'out/build-info').glob('*.json'):
        value=json.loads(path.read_text())
        if 'input' in value and 'output' in value: builds.append(value)
    targets=[]
    for transaction in broadcast['transactions']:
        if transaction['transactionType']=='CREATE':
            targets.append((transaction['contractAddress'], transaction['contractName'], transaction['transaction']['input']))
    targets.append((manifest['pair'],'UniswapV2Pair',None))
    statuses={}
    for address,name,creation in targets:
        result=api('getsourcecode',contractaddress=address,address=address)
        if result.get('status')=='1' and isinstance(result.get('result'),list) and result['result'][0].get('SourceCode'):
            statuses[address.lower()]={'verified':True,'contractName':name};print(name,'already verified',flush=True);continue
        selected=None
        for build in builds:
            for filename,contracts in build['output'].get('contracts',{}).items():
                if name not in contracts:continue
                compiled=contracts[name]
                bytecode=compiled['evm']['bytecode']['object']
                if creation is not None and not creation.removeprefix('0x').startswith(bytecode):continue
                if creation is None:
                    artifact=json.loads((root/'out'/Path(filename).name/(name+'.json')).read_text())
                    if bytecode!=artifact['bytecode']['object'].removeprefix('0x'):continue
                selected=(build,filename,bytecode);break
            if selected:break
        if selected is None:raise RuntimeError('Exact compiler build missing for '+name)
        build,filename,bytecode=selected
        compiler_input=build['input']
        settings=compiler_input['settings']
        # Narrow output generation only; retain every original source and optimizer setting.
        compiler_input['settings']['outputSelection']={filename:{name:['abi','evm.bytecode.object','evm.deployedBytecode.object','metadata']}}
        encoded=json.dumps(compiler_input)
        output=root/'deployments'/('verification-'+name+'.json');output.write_text(encoded)
        args=creation.removeprefix('0x')[len(bytecode):] if creation else ''
        version=build.get('solcLongVersion') or build.get('solc_long_version') or '0.8.24+commit.e11b9ed9'
        submitted=api('verifysourcecode',contractaddress=address,sourceCode=encoded,codeformat='solidity-standard-json-input',contractname=filename+':'+name,compilerversion='v'+version.removeprefix('v'),optimizationUsed='1',runs=str(settings['optimizer']['runs']),constructorArguments=args,evmVersion=settings['evmVersion'])
        if submitted.get('status')!='1':
            if 'already' in str(submitted.get('result','')).lower():statuses[address.lower()]={'verified':True,'contractName':name};continue
            raise RuntimeError(name+': '+str(submitted.get('result')))
        guid=submitted['result'];print(name,'submitted',flush=True)
        for attempt in range(18):
            time.sleep(5)
            checked=api('checkverifystatus',guid=guid)
            message=str(checked.get('result',''))
            if checked.get('status')=='1' or 'Already Verified' in message:
                statuses[address.lower()]={'verified':True,'contractName':name};print(name,'verified',flush=True);break
            if 'Pending' not in message and 'queue' not in message:
                raise RuntimeError(name+': '+message)
        else:raise RuntimeError('Verification still pending: '+name)
    manifest['explorerVerification']={name:statuses[manifest[name].lower()] for name in ['multisig','oracle','stable','weth','factory','router','pool','sp','jp','pair']}
    manifest['allSourcesVerified']=all(value['verified'] for value in manifest['explorerVerification'].values())
    manifest_path.write_text(json.dumps(manifest,indent=2)+'\n')
    print('All 10 contract sources verified.',flush=True)

if __name__=='__main__':
    try:main()
    except Exception as error:raise SystemExit(str(error).replace(api_key,'[redacted]'))
