import { assertEquals } from "jsr:@std/assert@1.0.11";
import path from "node:path";

const src = `
    <svg
      xmlns="http://www.w3.org/2000/svg"
      width="24"
      height="24"
      viewBox="0 0 24 24"
      fill="none"
      stroke="currentColor"
      stroke-width="2"
      stroke-linecap="round"
      stroke-linejoin="round"
    >
      <path d="M3.5 13h6" />
      <path d="m2 16 4.5-9 4.5 9" />
      <path d="M18 7v9" />
      <path d="m14 12 4 4 4-4" />
    </svg>
  `;

const want =
  `<svg xmlns="http://www.w3.org/2000/svg"width="24"height="24"viewBox="0 0 24 24"fill="none"stroke="currentColor"stroke-width="2"stroke-linecap="round"stroke-linejoin="round"><path d="M3.5 13h6"/><path d="m2 16 4.5-9 4.5 9"/><path d="M18 7v9"/><path d="m14 12 4 4 4-4"/></svg>`;

function runScript(srcPath: string, destPath: string) {
  const currentDir = new URL(".", import.meta.url).pathname;
  const scriptPath = path.join(currentDir, "html.ts");

  const cmd = new Deno.Command("deno", {
    args: [
      "run",
      "-A",
      scriptPath,
      srcPath,
      destPath,
    ],
  });

  const out = cmd.outputSync();
  if (out.success === false) {
    throw new Error(new TextDecoder().decode(out.stderr));
  }
}

Deno.test("minify HTML", () => {
  const tempDir = Deno.makeTempDirSync();
  const srcPath = `${tempDir}/src.html`;
  const destPath = `${tempDir}/dest.html`;
  Deno.writeTextFileSync(srcPath, src);

  runScript(srcPath, destPath);

  const destCont = Deno.readTextFileSync(destPath);
  assertEquals(destCont, want);
  Deno.removeSync(tempDir, { recursive: true });
});
