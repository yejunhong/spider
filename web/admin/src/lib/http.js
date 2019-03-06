import axios from 'axios'

const prefix = 'http://127.0.0.1:4321'

class Http {

  async get(uri){
    const res = await axios.get(`${prefix}${uri}`)
    return this.recover(res)
  }
  
  async post(){
  
  }

  recover(res){
    return res.data.msg
  }

}

export default new Http()