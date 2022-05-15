import { writable, Writable } from "svelte/store";

import type { ActionResponse, Environment } from "../../../lib";

// this is the interface element/control will be interacting with
export interface Store {
  set_value(field: string, val: any): void;
  register_before_submit(fn: () => void): void;

  // query_nested(field: string): Promise<any>;
  // verify_nested(field: string, data: any): Promise<any>;
}

interface State {
  stageTitle?: string;
  final: boolean;
  flowState:
    | "NOT_LOADED"
    | "SPLASH_LOADED"
    | "STAGE_LOADED"
    | "STAGE_PROCESSING"
    | "FINISHED";
  fields: object[];
  data_sources: { [_: string]: any };
  message?: string;
  epoch: number;
  errors?: { [_: string]: string };
}

export class WizardManager {
  _env: Environment;
  _opaqueData?: string;
  _wizard_title?: string;
  _state: Writable<State>;
  _fieldsStore: FieldsStore;
  _exec_options?: any;

  constructor(env: Environment, opts?: any) {
    this._env = env;
    this._state = writable({
      data_sources: {},
      fields: [],
      final: false,
      flowState: "NOT_LOADED",
      epoch: 0,
    });

    this._exec_options = opts;
    this._state.subscribe((state) => console.log("STATE =>", state));

    this._fieldsStore = null;
  }

  init = async () => {
    const resp = await this._env.PreformAction("get_splash", {
      has_exec_data: !!this._exec_options,
    });

    if (!resp.status_ok) {
      console.warn("error getting wizard splash", resp);
      return;
    }

    console.log("INIT RESP", resp);
    this.applySplashFields(resp);
    if (resp.body["skip_splash"]) {
      this.splash_next();
    }
  };

  applySplashFields = (resp: ActionResponse) => {
    this._fieldsStore = new FieldsStore(this);

    this._wizard_title = resp.body["wizard_title"] || "";
    const fields = resp.body["fields"] || [];
    const message = resp.body["message"] || "";
    const data_sources = resp.body["data_sources"] || {};
    this._state.update((old) => ({
      ...old,
      fields,
      epoch: old.epoch + 1,
      message,
      data_sources,
      flowState: "SPLASH_LOADED",
    }));
  };

  splash_next = async () => {
    const values = this._fieldsStore._get_values();

    const resp = await this._env.PreformAction("run_start", {
      data: values,
      exec_options: this._exec_options,
    });
    if (!resp.status_ok) {
      console.warn("error starting from splash", resp);
      return;
    }

    if (resp.body["stage_started"]) {
      this.applyStageFields(resp);
    } else {
      this.applySplashFields(resp);
    }
  };

  applyStageFields = (resp: ActionResponse) => {
    this._fieldsStore = new FieldsStore(this);
    const fields = resp.body["fields"] || [];
    this._opaqueData = resp.body["odata"] || "";
    const stageTitle = resp.body["stage_title"];
    this._state.update((old) => ({
      ...old,
      epoch: old.epoch + 1,
      fields,
      stageTitle,
      flowState: "STAGE_LOADED",
    }));
  };

  stage_next = async () => {
    const values = this._fieldsStore._get_values();
    this._state.update((old) => ({ ...old, flowState: "STAGE_PROCESSING" }));

    const resp = await this._env.PreformAction("run_next", {
      data: values,
      odata: this._opaqueData,
    });
    if (!resp.status_ok) {
      console.warn("error going to next stage", resp);
      return;
    }

    console.log("@=>", resp);

    if (!resp.body["ok"]) {
      const errors = resp.body["errors"] || {};
      this._state.update((old) => ({
        ...old,
        errors,
        flowState: "STAGE_LOADED",
      }));
      return;
    }

    if (resp.body["final"]) {
      const message = resp.body["last_message"] || "";
      this._state.update((old) => ({ ...old, flowState: "FINISHED", message }));
      return;
    }

    this.applyStageFields(resp);
  };

  stage_back = async () => {};
}

export class FieldsStore implements Store {
  _manager: WizardManager;
  _values: { [_: string]: any };

  _eventHandlers: Array<() => void>;
  constructor(m: WizardManager) {
    this._manager = m;
    this._values = new Map();
    this._eventHandlers = new Array(0);
  }

  set_value = (field: string, val: any) => {
    this._values[field] = val;
  };

  query_nested = async (field: string): Promise<any> => {
    return null;
  };

  verify_nested = async (field: string, data: any): Promise<any> => {
    return null;
  };

  register_before_submit(fn: () => void) {
    this._eventHandlers.push(fn);
  }

  // private

  _apply_event() {
    this._eventHandlers.forEach((eh) => {
      eh();
    });
  }

  _get_values() {
    console.log("@====>", this);

    this._apply_event();
    return this._values;
  }
}
