/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      colors: {
        ink: {
          900: '#0b0e14',
          800: '#161b22',
          700: '#1f2937',
        },
        accent: {
          500: '#3b82f6',
          600: '#2563eb',
        },
      },
    },
  },
  plugins: [],
}
