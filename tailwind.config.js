/** @type {import('tailwindcss').Config} */
const plugin = require('tailwindcss/plugin');
const theme = require('tailwindcss/defaultTheme');

module.exports = {
  content: ['./templates/*.html', './static/js/**/*.js'],
  theme: {
    extend: {},
  },
  plugins: [
    plugin(function ({ addBase }) {
      addBase({
        'h1': {
          '@apply text-primary text-3xl font-bold font-mono p-10',
        },
        'h2': {
          '@apply text-primary text-2xl font-bold font-mono pt-20 pb-10',
        },
        'h4': {
          '@apply text-primary text-20 font-bold font-mono pt-10 pb-10',
        },
        'p': {
          '@apply text-base leading-6 font-sans text-gray-700 mb-4 text-left',
        },
        'h3': {
          '@apply text-primary text-20 font-bold font-mono pt-10 pb-10',
        },
        'code': {
          '@apply font-bold',
        },
      });
    }),
  ],
};

