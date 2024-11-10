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
        secondary: '#8f613c',
        lightText: '#8d8ba0'
      },
      boxShadow: {
        'inner-custom': 'inset 0px 4px 8px rgba(0, 0, 0, 0.2)',
      }
    },
  },
  plugins: [],
}

