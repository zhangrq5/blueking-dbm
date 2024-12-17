import type { DetailBase, DetailClusters } from '../common';

export interface SingleRenameDatabase extends DetailBase {
  clusters: DetailClusters;
  force: boolean;
  infos: {
    cluster_id: number;
    force: boolean;
    from_database: string;
    to_database: string;
  }[];
}