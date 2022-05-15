import JsonSingleSelect from "./json/single_select.svelte";
import JsonMultiSelect from "./json/multi_select.svelte";
import JsonSingleNested from "./json/single_nested.svelte";
import JsonMultiNested from "./json/multi_nested.svelte";
import JsonSingleInline from "./json/single_inline.svelte";
import JsonMultiInline from "./json/multi_inline.svelte";

import ImageMultiSelect from "./image/multi_select.svelte";
import ViewParagraph from "./view/paragraph.svelte";
import * as Elem from "../service/elem_types";
import BasicElement from "./basic/basic.svelte";

const AdvElements = {
  [Elem.JSON_MULTI_SELECT]: JsonMultiSelect,
  [Elem.JSON_MULTI_INLINE]: JsonMultiInline,
  [Elem.JSON_MULTI_NESTED]: JsonMultiNested,
  [Elem.JSON_SINGLE_SELECT]: JsonSingleSelect,
  [Elem.JSON_SINGLE_INLINE]: JsonSingleInline,
  [Elem.JSON_SINGLE_NESTED]: JsonSingleNested,
  [Elem.IMAGE_MULTI_SELECT]: ImageMultiSelect,
};

const ViewElements = {
  [Elem.VIEW_PARAGRAPH]: ViewParagraph,
};

export { BasicElement, AdvElements, ViewElements };
