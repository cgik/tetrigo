import Link from "next/link";

export default function Header() {
    return (
        <header className="flex pt-4 pb-4">
            <div className="flex-1 left-0 pl-4">
                <Link href="/">tetrigo</Link>
            </div>
            <div className="flex-2 w-1/2 space-x-4">
                <Link href="/game/1">Play</Link>
                <Link href="/leaderboard">Leaderboard</Link>
                <Link href="/news">News</Link>
            </div>
            <div className="flex-3 right-0 pr-4">
                <Link href="/login">Login</Link>
            </div>
        </header>
    );
}