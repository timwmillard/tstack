/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    '../static/public/**/*.{html,js}',
    '**/*.templ',
    'node_modules/preline/dist/*.js',
  ],
  theme: {
    extend: {},
  },
  plugins: [
      require('@tailwindcss/forms'),
      require('preline/plugin')
  ],
  darkMode: 'class',
}

