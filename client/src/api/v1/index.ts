import api from '../index';

export default class V1 {
  public static getArticleList(config: any) {
    const path = 'getArticleList';
    return api.axios(path, config);
  }

  public static AddMemory(config: any) {
    const path = 'Add';
    return api.axios(path, config);
  }

  public static fecthMemory(config: any) {
    const path = 'Reply';
    return api.axios(path, config);
  }

  public static DeleteMemoryByAnswer(config: any) {
    const path = 'Forget';
    return api.axios(path, config);
  }

  public static FecthMemoryCount() {
    const path = 'Status';
    return api.axios(path);
  }
}
