/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./pages/**/*.{js,ts,jsx,tsx,mdx}",
    "./components/**/*.{js,ts,jsx,tsx,mdx}",
    "./app/**/*.{js,ts,jsx,tsx,mdx}",
  ],
  theme: {
    colors: {
      purple: "#7C5DFA",
      "purple-light": "#9277FF",
      "blue-pale": "#373B53",
      blue: "#1E2139",
      "blue-medium": "#252945",
      "blue-light": "#7E88C3",
      white: "#F8F8FB",
      "gray-light": "#DFE3FA",
      "gray-medium": "#888EB0",
      dark: "#0C0E16",
      "dark-medium": "#141625",
      orange: "#EC5757",
      "orange-light": "#FF9797",
    },
    extend: {
      backgroundImage: {
        "gradient-radial": "radial-gradient(var(--tw-gradient-stops))",
        "gradient-conic":
          "conic-gradient(from 180deg at 50% 50%, var(--tw-gradient-stops))",
      },
    },
  },
  plugins: [],
};
