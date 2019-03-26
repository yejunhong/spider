class Element{
  
  public handle: any;
  public constructor(handle: any) {
    this.handle = handle;
  }

  public async Html(selector: string): Promise<any> {
    return await this.handle.$eval(selector, e => e.innerHTML);
  }

  public async Attr(selector: string, attrName: string): Promise<any> {
    return await this.handle.$eval(selector, (e, attrName) => e.getAttribute(attrName), attrName);
  }

}
export default Element;