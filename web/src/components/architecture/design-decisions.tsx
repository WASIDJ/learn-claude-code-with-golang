"use client";

import { useState } from "react";
import { motion, AnimatePresence } from "framer-motion";
import { useTranslations, useLocale } from "@/lib/i18n";
import { ChevronDown } from "lucide-react";
import { cn } from "@/lib/utils";

<<<<<<< HEAD
import v0Annotations from "@/data/annotations/v0.json";
import v1Annotations from "@/data/annotations/v1.json";
import v2Annotations from "@/data/annotations/v2.json";
import v3Annotations from "@/data/annotations/v3.json";
import v4Annotations from "@/data/annotations/v4.json";
import v5Annotations from "@/data/annotations/v5.json";
import v6Annotations from "@/data/annotations/v6.json";
import v7Annotations from "@/data/annotations/v7.json";
import v8aAnnotations from "@/data/annotations/v8a.json";
import v8bAnnotations from "@/data/annotations/v8b.json";
import v8cAnnotations from "@/data/annotations/v8c.json";
import v0_miniAnnotations from "@/data/annotations/v0_mini.json";
import v9Annotations from "@/data/annotations/v9.json";
=======
import s01Annotations from "@/data/annotations/s01.json";
import s02Annotations from "@/data/annotations/s02.json";
import s03Annotations from "@/data/annotations/s03.json";
import s04Annotations from "@/data/annotations/s04.json";
import s05Annotations from "@/data/annotations/s05.json";
import s06Annotations from "@/data/annotations/s06.json";
import s07Annotations from "@/data/annotations/s07.json";
import s08Annotations from "@/data/annotations/s08.json";
import s09Annotations from "@/data/annotations/s09.json";
import s10Annotations from "@/data/annotations/s10.json";
import s11Annotations from "@/data/annotations/s11.json";
import s12Annotations from "@/data/annotations/s12.json";
>>>>>>> upstream/main

interface Decision {
  id: string;
  title: string;
  description: string;
  alternatives: string;
  zh?: { title: string; description: string };
  ja?: { title: string; description: string };
}

interface AnnotationFile {
  version: string;
  decisions: Decision[];
}

const ANNOTATIONS: Record<string, AnnotationFile> = {
<<<<<<< HEAD
  v0_mini: v0_miniAnnotations as AnnotationFile,
  v0: v0Annotations as AnnotationFile,
  v1: v1Annotations as AnnotationFile,
  v2: v2Annotations as AnnotationFile,
  v3: v3Annotations as AnnotationFile,
  v4: v4Annotations as AnnotationFile,
  v5: v5Annotations as AnnotationFile,
  v6: v6Annotations as AnnotationFile,
  v7: v7Annotations as AnnotationFile,
  v8a: v8aAnnotations as AnnotationFile,
  v8b: v8bAnnotations as AnnotationFile,
  v8c: v8cAnnotations as AnnotationFile,
  v9: v9Annotations as AnnotationFile,
=======
  s01: s01Annotations as AnnotationFile,
  s02: s02Annotations as AnnotationFile,
  s03: s03Annotations as AnnotationFile,
  s04: s04Annotations as AnnotationFile,
  s05: s05Annotations as AnnotationFile,
  s06: s06Annotations as AnnotationFile,
  s07: s07Annotations as AnnotationFile,
  s08: s08Annotations as AnnotationFile,
  s09: s09Annotations as AnnotationFile,
  s10: s10Annotations as AnnotationFile,
  s11: s11Annotations as AnnotationFile,
  s12: s12Annotations as AnnotationFile,
>>>>>>> upstream/main
};

interface DesignDecisionsProps {
  version: string;
}

function DecisionCard({
  decision,
  locale,
}: {
  decision: Decision;
  locale: string;
}) {
  const [open, setOpen] = useState(false);
<<<<<<< HEAD
=======
  const t = useTranslations("version");
>>>>>>> upstream/main

  const localized =
    locale !== "en" ? (decision as unknown as Record<string, unknown>)[locale] as { title?: string; description?: string } | undefined : undefined;

  const title = localized?.title || decision.title;
  const description = localized?.description || decision.description;

  return (
    <div className="rounded-lg border border-zinc-200 bg-white dark:border-zinc-700 dark:bg-zinc-900">
      <button
        onClick={() => setOpen(!open)}
        className="flex w-full items-center justify-between px-4 py-3 text-left"
      >
        <span className="pr-4 text-sm font-semibold text-zinc-900 dark:text-white">
          {title}
        </span>
        <ChevronDown
          size={16}
          className={cn(
            "shrink-0 text-zinc-400 transition-transform duration-200",
            open && "rotate-180"
          )}
        />
      </button>

      <AnimatePresence>
        {open && (
          <motion.div
            initial={{ height: 0, opacity: 0 }}
            animate={{ height: "auto", opacity: 1 }}
            exit={{ height: 0, opacity: 0 }}
            transition={{ duration: 0.2 }}
            className="overflow-hidden"
          >
            <div className="border-t border-zinc-100 px-4 py-3 dark:border-zinc-800">
              <p className="text-sm leading-relaxed text-zinc-600 dark:text-zinc-300">
                {description}
              </p>

              {decision.alternatives && (
                <div className="mt-3">
                  <h4 className="text-xs font-medium uppercase tracking-wide text-zinc-400 dark:text-zinc-500">
<<<<<<< HEAD
                    Alternatives Considered
=======
                    {t("alternatives")}
>>>>>>> upstream/main
                  </h4>
                  <p className="mt-1 text-sm leading-relaxed text-zinc-500 dark:text-zinc-400">
                    {decision.alternatives}
                  </p>
                </div>
              )}
            </div>
          </motion.div>
        )}
      </AnimatePresence>
    </div>
  );
}

export function DesignDecisions({ version }: DesignDecisionsProps) {
  const t = useTranslations("version");
  const locale = useLocale();

  const annotations = ANNOTATIONS[version];
  if (!annotations || annotations.decisions.length === 0) {
    return null;
  }

  return (
    <div className="space-y-4">
      <h2 className="text-xl font-semibold">{t("design_decisions")}</h2>
      <div className="space-y-2">
        {annotations.decisions.map((decision, i) => (
          <motion.div
            key={decision.id}
            initial={{ opacity: 0, y: 10 }}
            animate={{ opacity: 1, y: 0 }}
            transition={{ delay: i * 0.05 }}
          >
            <DecisionCard decision={decision} locale={locale} />
          </motion.div>
        ))}
      </div>
    </div>
  );
}
