export class DeviceInfo {
    id: string;
    user: string;
    imageUrl: string;
    historyUrl: string;

    constructor(options: any) {
        this.id = options.id;
        this.user = options.user;
        this.imageUrl = options.imageUrl;
        this.historyUrl = options.historyUrl;
    }
}
