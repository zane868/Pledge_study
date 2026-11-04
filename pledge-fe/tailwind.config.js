/** @type {import('tailwindcss').Config} */
export default {
  content: ['./index.html', './src/**/*.{ts,tsx}'],
  theme: {
    extend: {
      colors: {
        ink: {
          900: '#0b0e14',
          800: '#12161f',
          700: '#1a1f2b',
          600: '#252b39',
          500: '#39415280',
        },
        accent: { 500: '#3b82f6', 600: '#2563eb' },
        positive: '#16a34a',
        negative: '#dc2626',
        warn: '#d97706',
      },
    },
  },
  plugins: [],
};
