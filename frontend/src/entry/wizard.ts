import WizardApp from "./wizard/index.svelte"
import type { FactoryOptions } from "../lib";
import { registerExecLoaderFactory } from "../lib";



registerExecLoaderFactory("simplewizard.main", (opts: FactoryOptions) => {

    const __simple_wizard_app__ = new WizardApp({
        target: opts.target,
        props: {
            env: opts.env,
            options: opts.payload,
        }
    })
})