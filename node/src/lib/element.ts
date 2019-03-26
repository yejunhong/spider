class Element{
  
  public handle: any;
  public constructor(handle: any) {
    this.handle = handle;
  }

  /**
   * 
   * 根据选择器 获取html text
   * @param selector 选择器
   * @return Promise<string>
   */
  public async Html(selector: string): Promise<string> {
    return await this.handle.$eval(selector, e => e.innerHTML);
  }

  /**
   * 
   * 根据选择器，获取属性text
   * @param selector 选择器
   * @param attrName 属性名称
   * @return Promise<string>
   */
  public async Attr(selector: string, attrName: string): Promise<string> {
    return await this.handle.$eval(selector, (e, attrName) => e.getAttribute(attrName), attrName);
  }

  /**
   * 
   * 检测选择器是否存在
   * @param selector 选择器
   * @return Promise<boolean>
   */
  public async IsExist(selector: string): Promise<boolean> {
    const e = await this.handle.$eval(selector)
    if (e != null) {
      return true;
    }
    return false;
  }

}
export default Element;