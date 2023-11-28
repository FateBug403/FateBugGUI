<template>
  <div ref="editorRef" class="editor-main"></div>
</template>

<script lang="ts">
export default {
  name: 'pocDetails',
}
</script>

<script lang="ts" setup>
import { basicSetup, EditorView } from "codemirror";
import { EditorState } from "@codemirror/state";
import { oneDarkTheme } from "@codemirror/theme-one-dark";
// import { javascript } from "@codemirror/lang-javascript";
import { StreamLanguage,HighlightStyle,syntaxHighlighting  } from '@codemirror/language';
import { json } from '@codemirror/lang-json'
import { yaml } from '@codemirror/legacy-modes/mode/yaml';
import {tags} from "@lezer/highlight"

const props = defineProps({
  fatherPocString: {
    type: Object,
    default:{context:''}
  },
})

import { onMounted, ref } from "vue";
const editorRef = ref();
const editorView = ref();

const initEditor = () => {
  if (typeof editorView.value !== "undefined") {
    editorView.value.destroy();
  }

  const myHighlightStyle = HighlightStyle.define([
    {tag: tags.keyword, color: "#3fba2c"},
  ])

  const startState = EditorState.create({
    doc:props.fatherPocString.context,
    extensions: [basicSetup,oneDarkTheme,syntaxHighlighting(myHighlightStyle),StreamLanguage.define(yaml),json()],
  });

  if (editorRef.value) {
    editorView.value = new EditorView({
      state: startState,
      parent: editorRef.value,
    });
  }

};
// onMounted生命周期可以保证读取到dom元素
onMounted(() => {
  initEditor();
});

</script>

<style lang="scss" scoped>

//.editor-main{
//  width: 1200px;
//  height: 800px;
//}

</style>
