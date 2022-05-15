import { FactoryOptions, registerExecLoaderFactory } from "../lib/core/engine/plug/plug";
import SimpleWizardApp from "./simplewizard/index.svelte"


registerExecLoaderFactory("simplewizard.main", (opts: FactoryOptions) => {

    const __simple_wizard_app__ = new SimpleWizardApp({
        target: opts.target,
        props: {
            env: opts.env,
            options: opts.payload,
        }
    })
})