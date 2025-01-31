#!/usr/bin/env -S deno run -A

// @ts-types="npm:@types/html-minifier@4.0.5"
import { minify } from "npm:html-minifier@4.0.0";

if (Deno.args.length !== 2) {
  console.error("Usage: html.ts <source-file> <output-file>");
  console.log("The source file and output file can be the same.");
  Deno.exit(1);
}

const srcPath = Deno.args[0];
const srcCont = Deno.readTextFileSync(srcPath);

const destPath = Deno.args[1];
const destCont = minify(srcCont, {
  "collapseWhitespace": true,
  "removeTagWhitespace": true,
});
Deno.writeTextFileSync(destPath, destCont);

console.log("Minified HTML written to", destPath);
