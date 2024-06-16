import { Inter } from "next/font/google";
import PlausibleProvider from "next-plausible";
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
        <PlausibleProvider domain="tetrigo.ris.gg" />
      </head>
      <body className={inter.className}>{children}</body>
    </html>
  );
}
