import api from '../index'

export default class V1 {
  public static getArticleList(config: any) {
    let path = 'getArticleList'

    return api.axios(path, config)
  }

  public static AddMemory(config: any) {
    let path = 'Add'
    return api.axios(path, config)
  }

  public static fecthMemory(config: any) {
    let path = 'Reply'
    return api.axios(path, config)
  }
}
