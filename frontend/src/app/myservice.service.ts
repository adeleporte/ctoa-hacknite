import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable, of } from 'rxjs';
import { Ctoa } from './models';

@Injectable({
  providedIn: 'root'
})
export class MyserviceService {

  constructor(private http: HttpClient) { }

  getCtoas(): Observable<Ctoa[]> {
    return this.http.get<Ctoa[]>("/api/ctoa")
  }

}
