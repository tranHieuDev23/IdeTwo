import { Component, OnDestroy } from '@angular/core';
import { Subscription } from 'rxjs';
import {
  NotificationService,
  Notification,
} from './services/notification/notification.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss'],
})
export class AppComponent implements OnDestroy {
  public notificationList: Notification[] = [];

  private subscription: Subscription;

  constructor(readonly notificationService: NotificationService) {
    this.subscription =
      this.notificationService.notificationListChange.subscribe((value) => {
        this.notificationList = value;
      });
  }

  ngOnDestroy(): void {
    this.subscription.unsubscribe();
  }
}
