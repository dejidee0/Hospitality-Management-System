/** @type {import('tailwindcss').Config} */
export default {
  content: [
     "./index.html",
    "./src/**/*.{js,jsx,ts,tsx}",
  ],
  theme: {
    extend: {
      colors: {
        'primary-purple': '#5627ff',
        'auth-bg-color': '#fafafa',
        'light-background': '#ffffff',
        'default-grey': '#fafafa',
      },
    },
  },
  plugins: [],
}

