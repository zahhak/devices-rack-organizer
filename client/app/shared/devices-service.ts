import { Injectable, NgZone } from "@angular/core";
import { Observable, throwError } from "rxjs";
import { catchError } from "rxjs/operators";
import { DeviceInfo } from "~/shared/device.model";
import { HttpClient } from "~/shared/http-client";

@Injectable()
export class DevicesService {

    private serverUrl = "https://dp7o5mvps0.execute-api.eu-west-1.amazonaws.com";

    private get devicesEndpoint(): string {
        return `${this.serverUrl}/dev/device`;
    }

    private getDeviceActionEndpoint(deviceId): string {
        return `${this.serverUrl}/test/device/${deviceId}`;
    }

    constructor() { }

    public updateDeviceInfo(deviceId: string, user: string): Promise<void> {
        return HttpClient.call(this.getDeviceActionEndpoint(deviceId), "PUT", user);
    }

    public getAllDevices(user: string): Promise<DeviceInfo[]> {
        return HttpClient.call(this.devicesEndpoint, "GET", user)
            .then(res => res.content);
    }
}
