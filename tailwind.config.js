/** @type {import('tailwindcss').Config} */
const plugin = require('tailwindcss/plugin')
const theme = require("tailwindcss/defaultTheme");


module.exports = {
  content: [    "./templates/*.html", // Include HTML templates
    "./static/js/**/*.js",  ],
  theme: {
    extend: {},
  },
  plugins: [
  plugin(function({ addBase, config }) {
    addBase({
      'h1': {
        fontSize: config('theme.fontSize.3xl'),
        fontWeight: config('theme.fontWeight.bold'), // Adjusted fontWeight
        fontFamily: 'font-mono, monospace', // Adjusted fontFamily
        color: config('theme.colors.primary'),
        padding: '10px',
      },
      'h2': {
        fontSize: config('theme.fontSize.2xl'),
        fontWeight: config('theme.fontWeight.bold'), // Adjusted fontWeight
        fontFamily: 'font-mono, monospace', // Adjusted fontFamily
        color: config('theme.colors.primary'),
        paddingTop: '20px',
        paddingBottom: '10px'
      },
      'p': {
        fontSize: '16px', // Adjust font size
        lineHeight: '1.6', // Adjust line height for readability
        fontFamily: 'ui-sans-serif, system-ui, -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, "Helvetica Neue", Arial, "Noto Sans", sans-serif, "Apple Color Emoji", "Segoe UI Emoji", "Segoe UI Symbol", "Noto Color Emoji"', // Specify a default font family
        color: '#333', // Set the text color
        marginBottom: '10px', // Add some bottom margin for spacing between paragraphs
        textAlign: 'left', // Align text to the left
        // You can add more styles here if needed
      },

      'h3': { fontSize: config('theme.fontSize.lg') },
    })
  })
],
}

