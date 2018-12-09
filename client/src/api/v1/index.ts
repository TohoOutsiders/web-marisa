import api from '../index'

export default class V1 {
  public static getArticleList(config: any) {
    let path = 'getArticleList'

    return api.axios(path, config)
  }
}
