import { FactoryOptions, registerExecLoaderFactory } from "../lib/core/engine/plug/plug";
import DashedApp from "./dashed/index.svelte"


registerExecLoaderFactory("simpledash.main", (opts: FactoryOptions) => {
    new DashedApp({
        target: document.getElementById("plugroot"), // opts.target,
        props: {
            env: opts.env,
        }
    })
})