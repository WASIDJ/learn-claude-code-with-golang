import type { Metadata } from "next";
import { I18nProvider } from "@/lib/i18n";
import { Header } from "@/components/layout/header";
<<<<<<< HEAD
import "../globals.css";

export const metadata: Metadata = {
  title: "Learn Claude Code",
  description: "Build an AI coding agent from scratch, one concept at a time",
};

const locales = ["en", "zh", "ja"];
=======
import en from "@/i18n/messages/en.json";
import zh from "@/i18n/messages/zh.json";
import ja from "@/i18n/messages/ja.json";
import "../globals.css";

const locales = ["en", "zh", "ja"];
const metaMessages: Record<string, typeof en> = { en, zh, ja };
>>>>>>> upstream/main

export function generateStaticParams() {
  return locales.map((locale) => ({ locale }));
}

<<<<<<< HEAD
=======
export async function generateMetadata({
  params,
}: {
  params: Promise<{ locale: string }>;
}): Promise<Metadata> {
  const { locale } = await params;
  const messages = metaMessages[locale] || metaMessages.en;
  return {
    title: messages.meta?.title || "Learn Claude Code",
    description: messages.meta?.description || "Build an AI coding agent from scratch, one concept at a time",
  };
}

>>>>>>> upstream/main
export default async function RootLayout({
  children,
  params,
}: {
  children: React.ReactNode;
  params: Promise<{ locale: string }>;
}) {
  const { locale } = await params;

  return (
    <html lang={locale} suppressHydrationWarning>
<<<<<<< HEAD
=======
      <head>
        <script dangerouslySetInnerHTML={{ __html: `
          (function() {
            var theme = localStorage.getItem('theme');
            if (theme === 'dark' || (!theme && window.matchMedia('(prefers-color-scheme: dark)').matches)) {
              document.documentElement.classList.add('dark');
            }
          })();
        `}} />
      </head>
>>>>>>> upstream/main
      <body className="min-h-screen bg-[var(--color-bg)] text-[var(--color-text)] antialiased">
        <I18nProvider locale={locale}>
          <Header />
          <main className="mx-auto max-w-7xl px-4 py-8 sm:px-6 lg:px-8">
            {children}
          </main>
        </I18nProvider>
      </body>
    </html>
  );
}
