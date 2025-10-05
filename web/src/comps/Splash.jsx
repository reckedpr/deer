import { useState } from 'react';

const messages = [
  "fully open source !",
  "100% clanker free",
  "written by a deer (I swear)",
  "powered by astro.js",
  "compiled with love",
  "beep boop... kinda shy about it",
  "awaiting your request nya~",
  "give me head(ers)",
  "rm: use --no-preserve-root to override this failsafe",
  "aGVsbG8gY3V0aWU=",
  "import machine; machine.learn()",
  "git push --force hehe",
  "psitrance !!!",
  "if err != nil",
  "meow",
  "no dark mode 4 u"
];

export default function SplashFooter() {
  const [motd] = useState(() => messages[Math.floor(Math.random() * messages.length)]);
  return (
    <footer className="fixed bottom-0 w-full p-2 text-center text-sm sm:text-base">
      💌 <span className="text-slate-900/60">{motd}</span>
    </footer>
  );
}
