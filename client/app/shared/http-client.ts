import * as http from "tns-core-modules/http";
import { ImageSource } from "tns-core-modules/image-source/image-source";
import { HttpResponse } from "tns-core-modules/http";

export class HttpClient {
    public static call(url: string, method: "GET" | "PUT", token: string): Promise<http.HttpResponse> {
        const headers = token ?
            {
                Authorization: `Bearer ${token}`
            } : null;

        console.log("ABOUT TO CALL ", {
            url,
            method,
            headers
        });
        return http.request({
            url,
            method,
            headers
        });
    }

    public static getImage(url: string): Promise<ImageSource> {
        return http.getImage(url);
    }
}