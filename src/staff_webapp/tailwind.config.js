/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./src/**/*.{html,ts}",
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
      boxShadow: {
        'inner-custom': 'inset 0px 4px 8px rgba(0, 0, 0, 0.2)',
      }
    },
  },
  plugins: [],
}

