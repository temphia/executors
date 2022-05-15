import type { ApiManager } from "../../services"
import type { EngineAPI } from "../api/impl"
import {
    MODE_IFRAME,
    MODE_RAW_DOM,
    MODE_SUB_ORIGIN
} from "../core/engine/loader"

import { initFactory } from "../core/engine/registry"
import { generateId } from "../utils"
import { iframeTemplateBuild } from "./entry/iframe_build"

interface LauncherOptions {
    mode: string
    plug: string
    agent: string
    target: HTMLElement
    manager: ApiManager
}


export class PlugLauncher {
    mode: string
    token: string
    target: HTMLElement
    parent_secret: string
    plug: string
    agent: string
    base_url: string

    manager: ApiManager

    engine_api: EngineAPI

    private iframe?: HTMLIFrameElement
    // private sockdRooms: string[]

    constructor(opts: LauncherOptions) {
        this.mode = opts.mode
        this.target = opts.target
        this.plug = opts.plug
        this.agent = opts.agent
        this.manager = opts.manager
        this.parent_secret = generateId()
        initFactory() // fixme => if we are not dynamically bindings plugs this is not needed 
    }

    mount = () => {

        switch (this.mode) {
            case MODE_IFRAME:
                this.mountIframe()
                break;
            case MODE_RAW_DOM:
                this.mountRawDom()
                break
            case MODE_SUB_ORIGIN:
                this.mountSuborigin()
            default:
                console.warn("Mode not supported")
                break;
        }
    }

    private async mountIframe() {
        this.iframe = document.createElement("iframe");

        await this.loadEngine()

        const { data, status } = await this.engine_api.launcher_json(this.plug, this.agent, {})
        if (status !== 200) {
            console.warn("err loading", data)
            return
        }

        this.base_url = data["base_url"]

        console.log("__________DATA______________", data)

        const src =  iframeTemplateBuild({
            agent: this.agent,
            plug: this.plug,
            base_url: data["base_url"],
            entry_name: data["entry"],
            exec_loader: data["exec_loader"],
            js_plug_script: data["js_plug_script"],
            style_file: data["style"],
            token: data["token"] || "",
            ext_scripts: data["ext_scripts"],
            parent_secret: this.parent_secret
        })

        this.iframe.setAttribute("srcdoc", src);
        this.iframe.style.height = "100%"
        this.iframe.style.width = "100%"

        this.target.appendChild(this.iframe)
        window.addEventListener('message', this.on_message)
    }


    on_message = async (e) => {
        const decoded = JSON.parse(e.data);
        if (decoded.parent_secret !== this.parent_secret) {
            console.log("wrong parent token")
            return
        }
        console.log("ON_MESSAGE@PARENT", decoded)
    }




    unmount = () => {
        window.removeEventListener("message", this.on_message)
        this.target.removeChild(this.iframe)
        // fixme unregister rooms
    }

    private mountRawDom() {
        /*
            every js_app will have to call 
            window.RegisterPlugConstructor(name_of_plug, constructor)     
        */


    }


    private mountSuborigin() {


    }

    private loadEngine = async () => {
        if (this.engine_api) {
            return
        }
        
        this.engine_api = await this.manager.get_engine_api()
    }




}


// re export
export {
    MODE_IFRAME,
    MODE_RAW_DOM,
    MODE_SUB_ORIGIN,
}

