import { FactoryOptions, registerExecLoaderFactory } from "../lib/core/engine/plug/plug";
import SimpleDashApp from "./simpledash/index.svelte"


registerExecLoaderFactory("simpledash.main", (opts: FactoryOptions) => {
    new SimpleDashApp({
        target: document.getElementById("plugroot"), // opts.target,
        props: {
            env: opts.env,
        }
    })
})