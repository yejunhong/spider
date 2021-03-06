import axios from 'axios'

const prefix = 'http://127.0.0.1:4321'

class Http {

  async get(uri, param){
    const res = await axios.get(`${prefix}${uri}`, {params: param})
    return this.recover(res)
  }
  
  async post(uri, data){
    const res = await axios.post(`${prefix}${uri}`, data)
    return res.data
  }

  recover(res){
    return res.data.msg
  }

}

export default new Http()