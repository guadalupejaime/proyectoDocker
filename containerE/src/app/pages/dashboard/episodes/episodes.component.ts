import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-episodes',
  templateUrl: './episodes.component.html',
  styleUrls: ['./episodes.component.scss']
})
export class EpisodesComponent implements OnInit {
  status = 'load';
  constructor() { }

  ngOnInit(): void {
    setTimeout(() => {
      this.status = 'ok';
    }, 800);
  }

}
