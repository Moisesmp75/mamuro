<template>
  <h1 class="inset-0 text-5xl font-medium bg-violet-900 text-white underline decoration-double pt-5 pb-5 pl-3 pr-3">Mamuro Mail</h1>
  <div class=" pt-5 pb-5 pl-10 pr-10">
    <div class="flex">
      <Input v-model="this.query" />
      <Button textButton="Aceptar" :onClick="search_data" />
    </div>
    <DashBoard :data="this.mails" :query="this.query"/>
  </div>
</template>
<script>

import { MailServices } from '../services/mail-service';

import  Input   from '../components/Input.vue'
import Button from "../components/Button.vue"
import DashBoard from '../components/DashBoard.vue';

export default {
    
    components: { Button, Input, DashBoard },
    created() { 
      this.mailService = new MailServices()
    },
    async mounted() {
      await this.search_data()
    },
    data() {
      return {
        query: "inspected",
        pages: 10,
        mails: []
      }
    },
    methods: {
      async search_data() {
        const request = { query: this.query, pages: this.pages }
        const { resource } = await this.mailService.search_data(request)
        this.setMails(resource)
      },
      setMails(newMails) {
        this.mails = [...newMails]
      },
      created() {
        this.mails = [""]
      }
    }
}
</script>
<style>
  
</style>