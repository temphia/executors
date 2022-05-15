import { FactoryOptions, registerExecLoaderFactory } from "../lib/core/engine/plug/plug";
import WizardApp from "./wizard/index.svelte"


registerExecLoaderFactory("simplewizard.main", (opts: FactoryOptions) => {

    const __simple_wizard_app__ = new WizardApp({
        target: opts.target,
        props: {
            env: opts.env,
            options: opts.payload,
        }
    })
})