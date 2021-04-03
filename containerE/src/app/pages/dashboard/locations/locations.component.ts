import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-locations',
  templateUrl: './locations.component.html',
  styleUrls: ['./locations.component.scss']
})
export class LocationsComponent implements OnInit {
  status = 'load';

  constructor() {
  }

  ngOnInit(): void {
    setTimeout(() => {
      this.status = 'ok';
    }, 800);
  }

}
