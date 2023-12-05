
const url = "http://localhost:3000/api/v1"

export class MailServices {
  constructor() {
    this.requestOption = {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      }
    }
  }

  async search_data(query) {
    try {
      this.requestOption["body"] = JSON.stringify(query)
      const response = await fetch(`${url}/mail`, this.requestOption)
      if(!response.ok)
        throw new Error(response.status)
      const data = await response.json()
      return data
    } catch (error) {
      console.log(error)
      throw error
    }
  }
}