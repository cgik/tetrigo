import { Inter } from "next/font/google";
import Header from "@/components/Header";
import "./globals.css";

const inter = Inter({ subsets: ["latin"] });

export const metadata = {
    title: "tetrigo",
    description: "A tetris attack reimagined in the browser.",
};

export default function RootLayout({ children }) {
    return (
        <html lang="en">
        <body className={inter.className}>{children}</body>
        </html>
    );
}