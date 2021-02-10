import { Component, OnInit } from '@angular/core';
import { MyserviceService } from './myservice.service';
import { Ctoa } from './models';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit {
  title = 'ctoa-web';
  ctoas = [];

  constructor(public myService: MyserviceService) { }

  ngOnInit() {
    this.myService.getCtoas().subscribe(data => {console.log(data); this.ctoas = data} )
  }

  refresh() {
    this.myService.getCtoas().subscribe(data => {console.log(data); this.ctoas = data} )
  }
}
