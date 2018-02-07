import React from "react";

import classes from './SidebarList.css';
import NavigationItem from './NavigationItem/NavigationItem';

const sidebarList = () => {
    const items = [
        {
            link: "/",
            icon: "fa-home",
            label: "Dashboard"
        },
        {
            link: "/upload-files",
            icon: "fa-cloud-upload",
            label: "Upload Files"
        }
    ];

    return <aside className={classes.SidebarList}>
        <ul>
            {items.map((item, index) => (<li key={index}>
                <NavigationItem
                    id={index}
                    link={item.link}
                >
                    <i className={"fa "+item.icon} />
                    {item.label}
                </NavigationItem>
            </li>))}
        </ul>
    </aside>
};

export default sidebarList;