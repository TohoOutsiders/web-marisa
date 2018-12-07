import api from '../index'

export default class V1 {
  public static getIndex(config: any) {
    let path = 'getIndex'

    return api.axios(path, config)
  }
}
