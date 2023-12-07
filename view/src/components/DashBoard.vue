<template>
  <div class="p-4">
    <h3 class="text-2xl font-bold mb-4">Dashboard</h3>
    <div class="overflow-x-auto">
      <table class="min-w-full table-fixed">
        <thead class="">
          <tr>
            <th class="px-4 py-2">MessageId</th>
            <th class="px-4 py-2">From</th>
            <th class="px-4 py-2">To</th>
            <th class="px-4 py-2">Content</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(mail, i) in mails" :key="i" class="border-t align-top">
            <td class="px-4 py-2">{{ mail.message_id }}</td>
            <td class="px-4 py-2">{{ mail.from }}</td>
            <td class="px-4 py-2">
              <ul class="list-inside list-none">
                <li v-for="(item, i) in mail.to" :key="i" class="whitespace-nowrap">{{ item }}</li>
              </ul>
            </td>
            <td class="px-4 py-2" v-html="formatContent(mail.content)">
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>
<script>
export default {
  props: {
    data: {
      type: Array,
      required: true
    },
    query: {
      type: String,
      required: true
    }
  },
  data() {
    return {
      mails: []
    }
  },
  created() {
    this.mails = this.data
  },
  updated() {
    this.mails = this.data
  },
  methods: {
    formatContent(content) {
      const queryPosition = content.toLowerCase().indexOf(this.query.toLowerCase());

      const paragraphs = content.split('\n\n'); // Dividir por párrafos
      const formattedContent = paragraphs.map(paragraph => `${paragraph}`).join('');

      if(queryPosition === -1) {

      }

      const shortContent = (queryPosition > 40 ? "..." : "") + formattedContent.slice(queryPosition < 40 ? 0 : queryPosition - 40, queryPosition + 85) + "..."

      const highlightedContent = shortContent.replace(this.query, `<span style="background-color: yellow; font-weight: bold;">${this.query}</span>`);
      return highlightedContent
    }
  }

}
</script>
<style></style>