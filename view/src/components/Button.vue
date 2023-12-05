<template>
  <button class="bg-violet-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded" v-on:click="onClick" >{{ textButton }}</button>
</template>
<script>
export default {
  props: ["textButton", "onClick"]
}
</script>
<script setup>

defineProps({
  textButton: { 
    type: String,
    required: true
  },
  onClick: {
    type: Function,
    required: true
  }
})

</script>

<style lang="">
  
</style>