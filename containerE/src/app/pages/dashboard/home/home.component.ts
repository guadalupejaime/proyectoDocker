import { Component, OnInit } from '@angular/core';
import { Character } from '@pages/dashboard/characters/characters.interface';
import { HttpService } from '@services/http.service';
import { Subject } from 'rxjs';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.scss']
})
export class HomeComponent implements OnInit {
  status = 'load';
  characters: Character[];
  totalCharacters: number;
  totalEpisodes: number;
  totalLocations: number;
  readyHome: Subject<boolean> = new Subject();
  contSubscriptors = 0;

  constructor(
    private sHttpService: HttpService
  ) { }

  ngOnInit(): void {
    this.getCharacters();
    this.getEpisodes();
    this.getLocations();

    this.readyHome.subscribe(
      (isComplete: boolean) => {
        if (isComplete) {
          this.contSubscriptors++;
          if (this.contSubscriptors === 3) {
            this.status = 'ok';
          }
        } else {
          this.status = 'error';
        }
      }
    );
  }

  /** get characters from api */
  getCharacters(): void {
    this.sHttpService.getCharacters().subscribe(
      (characters) => {
        this.totalCharacters = characters.total_found;
        let start = Math.floor((Math.random() * characters.results.length - 1) + 1);
        if (start + 6 > characters.results.length) {
          start = 0;
        }
        this.characters = characters.results.slice(start, start + 6);
        this.readyHome.next(true);
      },
      (err) => {
        this.readyHome.next(false);
      }
    );
  }

  /** get episodes from api */
  getEpisodes(): void {
    this.sHttpService.getEpisodes().subscribe(
      (episodes) => {
        this.totalEpisodes = episodes.total_found;
        this.readyHome.next(true);
      },
      (err) => {
        this.readyHome.next(false);
      }
    );
  }

  /** get locations from api */
  getLocations(): void {
    this.sHttpService.getLocations().subscribe(
      (locations) => {
        this.totalLocations = locations.total_found;
        this.readyHome.next(true);
      },
      (err) => {
        this.readyHome.next(false);
      }
    );
  }
}
