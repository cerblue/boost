import {useQuery} from "@apollo/react-hooks";
import {SealingPipelineQuery} from "./gql";
import React from "react";
import {humanFileSize} from "./util";
import {PageContainer, ShortDealLink} from "./Components";
import './SealingPipeline.css'
import {dateFormat} from "./util-date";
import moment from 'moment';
import {NavLink} from "react-router-dom";
import {CumulativeBarChart} from "./CumulativeBarChart";

export function SealingPipelinePage(props) {
    return <PageContainer icon={<SealingPipelineIcon />} title="Sealing Pipeline">
        <SealingPipelineContent />
    </PageContainer>
}

function SealingPipelineContent(props) {
    const {loading, error, data} = useQuery(SealingPipelineQuery, { pollInterval: 2000 })

    if (loading) {
        return <div>Loading...</div>
    }
    if (error) {
        return <div>Error: {error.message}</div>
    }

    const sealingPipeline = data.sealingpipeline

    return <div className="sealing-pipeline">
        <WaitDeals wdSectors={sealingPipeline.WaitDealsSectors} sdwdSectors={sealingPipeline.SnapDealsWaitDealsSectors} />
        <Sealing states={sealingPipeline.SectorStates} />
        <Workers workers={sealingPipeline.Workers} />
    </div>
}

function WaitDeals(props) {
    return <div className="wait-deals split-pane">
        <div className="wait-deals-type split-pane-half">
            <WaitDealsType sectors={props.wdSectors} name="Regular"/>
        </div>
        <div className="wait-deals-type split-pane-half">
            <WaitDealsType sectors={props.sdwdSectors} name="Snap Deals"/>
        </div>
    </div>
}

function WaitDealsType(props) {
    return <div>
        <h5>{props.name} Wait Deals</h5>
        { props.sectors.length ? (
            props.sectors.map(s => <WaitDealsSector sector={s}/>)
        ) : (
            <div className="no-deals">There are no sectors in the WaitDeals state</div>
        ) }
    </div>
}

function WaitDealsSector(props) {
    const sector = props.sector
    const free = sector.SectorSize - sector.Used
    const haveDeals = sector.Deals.length > 0

    var bars = [{
        className: 'free',
        amount: free,
    }]
    if (haveDeals) {
        bars = [{
            className: 'filled',
            amount: sector.Used,
        }].concat(bars)
    }

    return <div className="wait-deals-sector">
        <div className="sector-id">Sector {sector.SectorID+''}</div>
        <CumulativeBarChart bars={bars} unit="byte" />

        { haveDeals ? <WaitDealsSizes free={free} deals={sector.Deals} /> : (
            <div className="no-deals">There are no deals in this sector</div>
        )}
    </div>
}

function WaitDealsSizes(props) {
    return <table className="wait-deals-sizes">
        <tbody>
            {props.deals.map(deal => (
                <tr key={deal.ID}>
                    <td className="deal-id">
                        {deal.IsLegacy ? (
                            <NavLink to={"/legacy-deals/" + deal.ID}>
                                <div className="short-deal-id">{deal.ID.substring(0, 12) + '…'}</div>
                            </NavLink>
                        ) : (
                            <ShortDealLink id={deal.ID} />
                        )}
                    </td>
                    <td className="deal-size">{humanFileSize(deal.Size)}</td>
                </tr>
            ))}
            <tr key="free" className="free">
                <td className="deal-id">Free</td>
                <td className="deal-size">{humanFileSize(props.free)}</td>
            </tr>
        </tbody>
    </table>
}

function Sealing(props) {
    return (
        <div className="sealing split-pane section">
            <div className="sealing-type split-pane-half">
                <SealingType states={props.states} sectorType="Regular" />
            </div>
            <div className="sealing-type split-pane-half">
                <SealingType states={props.states} sectorType="Snap Deals" />
            </div>
        </div>
    )
}

function SealingType(props) {
    var states = []
    var errStates = []
    if (props.sectorType === 'Regular') {
        states = props.states.Regular
        errStates = props.states.RegularError
    } else {
        states = props.states.SnapDeals
        errStates = props.states.SnapDealsError
    }
    states = states.map(s => s).sort((a, b) => a.Order - b.Order)
    states = states.concat(errStates.map(s => s).sort((a, b) => a.Order - b.Order))

    const title = props.sectorType === 'Snap Deals' ? 'Snap Deals Sectors' : 'Regular Sectors'
    const className = props.sectorType.toLowerCase().replace(/ /g, '_')
    return <div className={"sealing-type-content " + className}>
        <h5>{title}</h5>

        <table className="sector-states">
            <tbody>
            {states.map(sec => (
                <tr key={sec.Key}>
                    <td className="color"><div className={sec.Key.replace(/ /g, '_')} /></td>
                    <td className="state">{sec.Key}</td>
                    <td className={"count " + (sec.Value ? '' : 'zero')}>{sec.Value}</td>
                </tr>
            ))}
            </tbody>
        </table>
    </div>
}

function Workers(props) {
    return <div className="workers">
        <h5>Workers</h5>
        {props.workers.length === 0 ? <div className="no-workers">No active jobs</div> : (
            <table>
                <tbody>
                <tr>
                    <th className="start">Start</th>
                    <th className="worker-id">ID</th>
                    <th className="stage">Stage</th>
                    <th className="sector">Sector</th>
                </tr>
                {props.workers.map(worker => (
                    <tr key={worker.ID}>
                        <td className="start">{moment(worker.Start).format(dateFormat)}</td>
                        <td className="worker-id">{worker.ID}</td>
                        <td className="stage">{worker.Stage}</td>
                        <td className="sector">{worker.Sector}</td>
                    </tr>
                ))}
                </tbody>
            </table>
        )}
    </div>
}

export function SealingPipelineMenuItem(props) {
    const {data} = useQuery(SealingPipelineQuery, {
        pollInterval: 5000,
        fetchPolicy: "network-only",
    })

    var total = 0
    if (data) {
        for (let sec of data.sealingpipeline.SectorStates.Regular) {
            total += Number(sec.Value)
        }
        for (let sec of data.sealingpipeline.SectorStates.RegularError) {
            total += Number(sec.Value)
        }
        for (let sec of data.sealingpipeline.SectorStates.SnapDeals) {
            total += Number(sec.Value)
        }
        for (let sec of data.sealingpipeline.SectorStates.SnapDealsError) {
            total += Number(sec.Value)
        }
    }

    return <NavLink key="sealing-pipeline" className="sidebar-item sidebar-item-deals" to="/sealing-pipeline">
        <span className="sidebar-icon">
            <SealingPipelineIcon />
        </span>
        <span className="sidebar-title">Sealing Pipeline</span>
        <div className="sidebar-item-excerpt">
          <span className="figure">{total}</span>
          <span className="label">Sector{total === 1 ? '' : 's'}</span>
        </div>
    </NavLink>
}

export function SealingPipelineIcon(props) {
    return <svg width="26" height="22" viewBox="0 0 26 22" fill="none" xmlns="http://www.w3.org/2000/svg">
        <rect x="1" y="1" width="24" height="6" rx="2" strokeWidth="2"/>
        <path
            d="M14 1a1 1 0 1 0-2 0h2zm-1.707 20.707a1 1 0 0 0 1.414 0l6.364-6.364a1 1 0 0 0-1.414-1.414L13 19.586l-5.657-5.657a1 1 0 0 0-1.414 1.414l6.364 6.364zM12 1v20h2V1h-2z"
            fill="currentColor"
            className="both"/>
    </svg>
}
