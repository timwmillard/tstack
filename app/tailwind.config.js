/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    '../static/public/**/*.{html,js}',
    './node_modules/flowbite/**/*.js',
    '**/*.templ'
  ],
  theme: {
    extend: {},
  },
  plugins: [
    require('flowbite/plugin')
  ],
}

