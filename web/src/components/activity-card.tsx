import {
  Button,
  Card,
  CardFooter,
  CardHeader,
  Image,
  Spacer,
} from "@heroui/react";

import { formatDate } from "@/utils/datetime.ts";

interface ActivityCardProps {
  activity: any;
}

export const ActivityCard = ({ activity }: ActivityCardProps) => {
  return (
    <Card key={activity.id} isFooterBlurred className="py-1">
      <CardHeader className="pb-0 pt-2 px-4 flex-col items-start">
        <h4 className="font-bold text-large">{activity.name}</h4>
        <p className="text-tiny uppercase font-bold">
          时间: {formatDate(activity.start_at)} - {formatDate(activity.end_at)}
        </p>
      </CardHeader>
      <Spacer y={2} />
      <Image
        removeWrapper
        alt="Card background"
        className="object-cover rounded-xl h-full"
        src={
          activity.avatar
            ? activity.avatar
            : "https://heroui.com/images/hero-card.jpeg"
        }
      />
      <CardFooter className="absolute bg-black/40 bottom-0 z-10 border-t-1 border-default-600 dark:border-default-100 h-1/5">
        <div className="flex flex-grow gap-2 items-center">
          <Image
            alt="activity icon"
            className="rounded-full w-10 h-11 bg-black"
            src={
              activity.creator.avatar
                ? activity.creator.avatar
                : "https://heroui.com/images/hero-card.jpeg"
            }
          />
          <div className="flex flex-col">
            <p className="text-tiny text-white/60">
              {activity.creator.nickname}
            </p>
          </div>
        </div>
        <Button radius="full" size="sm">
          查看
        </Button>
      </CardFooter>
    </Card>
  );
};
