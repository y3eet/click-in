import Image from "next/image";
import React from "react";
import { Button } from "./ui/button";

export type ClickableProps = {
  imageUrl: string;
  onClick?: React.MouseEventHandler<HTMLButtonElement>;
  alt?: string;
  disabled?: boolean;
  className?: string;
  size?: number; // px
};

export function Clickable({
  imageUrl,
  onClick,
  alt = "Clickable",
  disabled = false,
  className,
  size = 350,
}: ClickableProps) {
  return (
    <>
      <Button
        type="button"
        className={["clickable-btn", className].filter(Boolean).join(" ")}
        onClick={onClick}
        disabled={disabled}
        style={
          {
            "--btn-size": `${size}px`,
            maxWidth: size,
            maxHeight: size,
          } as React.CSSProperties
        }
        aria-label={alt}
      >
        <Image
          unoptimized
          className="clickable-img"
          src={imageUrl}
          alt={alt}
          draggable={false}
          width={1500}
          height={1500}
        />
      </Button>

      <style>{`
        .clickable-btn {
          display: inline-grid;
          place-items: center;
          padding: 0;
          border: none;

          /* responsive sizing - respects size prop */
          width: min(var(--btn-size, 350px), 100%);
          height: min(var(--btn-size, 350px), 100%);
          aspect-ratio: 1;

          /* more rounded */
          border-radius: 9999px;

          background: rgba(255, 255, 255, 0.07);
          cursor: pointer;
          user-select: none;
          -webkit-tap-highlight-color: transparent;

          box-shadow:
            0 10px 26px rgba(0,0,0,0.20),
            inset 0 0 0 1px rgba(255,255,255,0.14);

          overflow: hidden;
          transform: translateZ(0);
          will-change: transform, box-shadow, filter;

          /* snappier */
          transition:
            transform 90ms cubic-bezier(.2,.9,.2,1),
            box-shadow 110ms cubic-bezier(.2,.9,.2,1),
            filter 90ms cubic-bezier(.2,.9,.2,1);
          animation: clickable-pop 180ms cubic-bezier(.16,1,.3,1);
        }

        /* responsive breakpoints */
        @media (max-width: 640px) {
          .clickable-btn {
            width: min(var(--btn-size, 350px), 90vw);
            height: min(var(--btn-size, 350px), 90vw);
          }
        }

        @media (min-width: 641px) and (max-width: 1024px) {
          .clickable-btn {
            width: min(var(--btn-size, 350px), 80vw);
            height: min(var(--btn-size, 350px), 80vw);
          }
        }

        .clickable-img {
          width: 100%;
          height: 100%;
          object-fit: cover;
          border-radius: inherit;

          transform: scale(1.03);
          will-change: transform, filter;
          transition: transform 110ms cubic-bezier(.2,.9,.2,1), filter 110ms cubic-bezier(.2,.9,.2,1);
        }

        .clickable-btn:hover {
          transform: translateY(-1px) scale(1.02);
          box-shadow:
            0 14px 34px rgba(0,0,0,0.26),
            inset 0 0 0 1px rgba(255,255,255,0.18);
        }

        .clickable-btn:hover .clickable-img {
          transform: scale(1.08);
        }

        .clickable-btn:active {
          transform: translateY(0px) scale(0.96);
        }

        .clickable-btn:active .clickable-img {
          transform: scale(1.03);
          filter: saturate(1.05) contrast(1.03);
        }

        .clickable-btn:focus-visible {
          outline: none;
          box-shadow:
            0 0 0 3px rgba(99,102,241,0.55),
            0 16px 40px rgba(0,0,0,0.26),
            inset 0 0 0 1px rgba(255,255,255,0.18);
        }

        .clickable-btn:disabled {
          cursor: not-allowed;
          opacity: 0.55;
          transform: none;
          animation: none;
        }

        @keyframes clickable-pop {
          from { transform: scale(0.96); opacity: 0.001; }
          to   { transform: scale(1); opacity: 1; }
        }

        @media (prefers-reduced-motion: reduce) {
          .clickable-btn { transition: none; animation: none; }
          .clickable-img { transition: none; }
        }
      `}</style>
    </>
  );
}
