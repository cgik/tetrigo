import { Inter } from "next/font/google";
import Navigation from "@/components/Navigation";
import "./globals.css";

const inter = Inter({ subsets: ["latin"] });

export const metadata = {
  title: "tetrigo",
  description: "A tetris attack reimagined in the browser.",
};

export default function RootLayout({ children }) {
  return (
    <html data-theme="dark" lang="en">
      <head>
        <script
          defer
          data-domain="tetrigo.ris.gg"
          src="/modules/js/script.js"
        />
      </head>
      <body className={inter.className}>{children}</body>
    </html>
  );
}
