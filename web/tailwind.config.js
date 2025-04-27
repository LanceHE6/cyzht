import { heroui } from "@heroui/theme";

/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./index.html",
    "./src/layouts/**/*.{js,ts,jsx,tsx,mdx}",
    "./src/pages/**/*.{js,ts,jsx,tsx,mdx}",
    "./src/components/**/*.{js,ts,jsx,tsx,mdx}",
    "./node_modules/@heroui/theme/dist/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      screens: {
        xs: '480px',
      },
      gridTemplateColumns: {
        "16": 'repeat(16, minmax(0, 1fr))',
        "24": 'repeat(24, minmax(0, 1fr))',
        "15": 'repeat(15, minmax(0, 1fr))'
      },
      gridColumn: {
        "span-16": "span 16 / span 16",
        "span-17": "span 17 / span 17",
        "span-18": "span 18 / span 18",
        "span-19": "span 19 / span 19",
        "span-20": "span 20 / span 20",
        "span-21": "span 21 / span 21",
        "span-22": "span 22 / span 22",
        "span-23": "span 23 / span 23",
        "span-24": "span 24 / span 24",
        "span-15": "span 15 / span 15"
      }
    },
  },
  darkMode: "class",
  plugins: [heroui()],
};
