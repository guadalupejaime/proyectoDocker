import { Component, OnInit } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { PageEvent } from '@angular/material/paginator';
import { MatSnackBar } from '@angular/material/snack-bar';
import { MatTableDataSource } from '@angular/material/table';
import { Episode } from '@pages/dashboard/episodes/episodes.interface';
import { AddEpisodeComponent } from '@pages/dashboard/episodes/modals/add-episode/add-episode.component';
import { InfoEpisodeComponent } from '@pages/dashboard/episodes/modals/info-episode/info-episode.component';
import { TablePage } from '@services/http.interfaces';
import { HttpService } from '@services/http.service';

@Component({
  selector: 'app-episodes',
  templateUrl: './episodes.component.html',
  styleUrls: ['./episodes.component.scss']
})
export class EpisodesComponent implements OnInit {
  status = 'load';
  displayedColumns: string[] = ['name', 'air_date', 'episode', 'info'];
  totalEpisodes = 0;
  episodes: Episode[] = [];
  dataEpisodes = new MatTableDataSource<Episode>(this.episodes);
  sizePage = 20;
  paginatorSettings: TablePage;

  constructor(
    private sHttpService: HttpService,
    private sMatDialog: MatDialog,
    private sMatSnackBar: MatSnackBar
  ) { }

  ngOnInit(): void {
    this.paginatorSettings = {
      limit: this.sizePage,
      offset: 0
    };

    this.sHttpService.getEpisodes(this.paginatorSettings).subscribe(
      (episodes) => {
        this.totalEpisodes = episodes.total_found;
        this.dataEpisodes.data = episodes.results;
        this.status = 'ok';
      },
      (err) => {
        this.status = 'error';
      }
    );
  }

  /** update episodes table */
  updateEpisodes(): void {
    this.sHttpService.getEpisodes(this.paginatorSettings).subscribe(
      (episodes) => {
        this.totalEpisodes = episodes.total_found;
        this.dataEpisodes.data = episodes.results;
      },
      (err) => {
      }
    );
  }

  /** update paginator settings */
  updatePaginatorSettings(event: PageEvent): void {
    this.paginatorSettings = {
      limit: event.pageSize,
      offset: event.pageIndex * event.pageSize
    };
    this.updateEpisodes();
  }

  /** search on data table */
  applyFilter(event: Event): void {
    const filterValue = (event.target as HTMLInputElement).value;
    this.dataEpisodes.filter = filterValue.trim().toLowerCase();
  }

  /** open a modal wt episode information selected */
  infoEpisode(id: number): void {
    const modal = this.sMatDialog.open(InfoEpisodeComponent, {
      maxHeight: '99%',
      maxWidth: '99%',
      disableClose: true,
      panelClass: 'myModal',
      backdropClass: 'myBackdrop',
      data: { idEpisode: id }
    });
  }


  /** open a  modal to add a new episode */
  addEpisode(): void {
    const modal = this.sMatDialog.open(AddEpisodeComponent, {
      maxHeight: '99%',
      maxWidth: '99%',
      disableClose: true,
      panelClass: 'myModal',
      backdropClass: 'myBackdrop',
    });

    modal.afterClosed().subscribe(
      (result) => {
        if (result) {
          this.openNotification('Added new episode, success!');
          this.updateEpisodes();
          return;
        } else if (result === 'error') {
          this.openNotification('Something went wrong!');
        }
      });
  }

  openNotification(message: string): void {
    this.sMatSnackBar.open(message, 'End now', {
      duration: 1500,
      horizontalPosition: 'right',
      verticalPosition: 'top',
    });
  }

}
