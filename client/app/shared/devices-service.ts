import { Injectable, NgZone } from "@angular/core";
import { Observable, throwError } from "rxjs";
import { catchError } from "rxjs/operators";
import { DeviceInfo } from "~/shared/device.model";
import { DeviceStatus } from "~/shared/deviceStatus.model";
import { HttpClient } from "~/shared/http-client";

@Injectable()
export class DevicesService {

    private _devices: DeviceInfo[];
    private serverUrl = "https://dp7o5mvps0.execute-api.eu-west-1.amazonaws.com/dev/device";
    private serverUrl2 = "https://b04hmprozf.execute-api.eu-west-1.amazonaws.com/test/device";

    private get devicesEndpoint(): string {
        return `${this.serverUrl}`;
    }

    private getDeviceActionEndpoint(deviceId): string {
        return `${this.serverUrl2}/${deviceId}`;
    }

    constructor() { }

    public updateDeviceInfo(deviceId: string, user: string): Promise<DeviceStatus> {
        return HttpClient.call(this.getDeviceActionEndpoint(deviceId), "PUT", user)
            .then(res => res.content.toJSON());
    }

    public getAllDevices(user: string): Promise<DeviceInfo[]> {
        return HttpClient.call(this.devicesEndpoint, "GET", user)
            .then(res => {
                this._devices = res.content.toJSON();
                return this._devices;
            });
    }

    public getDeviceInfo(user: string, deviceId: string): Promise<DeviceInfo> {
        const getDevicesPromise = (!this._devices || this._devices.length === 0) ?
            this.getAllDevices(user) : Promise.resolve(this._devices);

        return getDevicesPromise.then(res => {
            const devicesFound = this._devices.filter(d => d.id === deviceId);
            return devicesFound && devicesFound[0];
        });
    }
}
