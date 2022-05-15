
import Example from "./example.svelte"


const Controls = new Map()



Controls.set("example", Example)



export interface IRowEditor {
    RegisterBeforeSave(field: string, callback: () => [boolean, any]): void
    OnChange(field: string, value: any): void
}

console.log(Example)
