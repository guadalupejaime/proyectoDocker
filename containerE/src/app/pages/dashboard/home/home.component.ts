import { Component, OnInit } from '@angular/core';
import { Character } from '@pages/dashboard/characters/characters.interface';
import { HttpService } from '@services/http.service';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.scss']
})
export class HomeComponent implements OnInit {
  characters: Character[];

  constructor(
    private sHttpService: HttpService
  ) { }

  ngOnInit(): void {
    this.sHttpService.getCharacters().subscribe(
      (characters) => {
        let start = Math.floor((Math.random() * characters.results.length - 1) + 1);
        if (start + 6 > characters.results.length) {
          start = 0;
        }
        this.characters = characters.results.slice(start, start + 6);
        // console.log(characters.results);
        console.log(this.characters);
      },
      (err) => {

      }
    );
  }
}
