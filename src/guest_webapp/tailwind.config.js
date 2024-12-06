/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./src/**/*.{js,jsx,ts,tsx}",
  ],
  theme: {
    extend: {
      colors: {
        primary: '#ecd8c8',
        lightPrimary: '#faf2eb',
        secondary: '#C9B8AA',
        lightText: '#8d8ba0',
        brown: '#8f613c',
        lightGreen: '#07bc0c',
        lightRed: '#e74c3c'
      },
      fontFamily: {
        futura: ['Futura', 'Helvetica Neue', 'sans-serif'],
      }
    },
  },
  plugins: [],
}

