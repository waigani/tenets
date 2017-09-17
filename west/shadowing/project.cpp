PCSSCall CCSSInterfaceImp::makeCallInternal(CLine *line, CLineGroup *lineGroup, CAgency *agency, PIncident incident, UseIncidentFor incidentUse)
{
    if( cssCall != NULL )
    {
      {
        // try to match inIncident to a centrally created site call
        CThreadCSLock _lock(mMutex); // insure proper order of mutexes
        CThreadCSLock _lock1(MUTEX(cssCall));
        if( matchableByCentralUCId(cssCall) )
        {
          for( SiteCallMap::reverse_iterator incidentIt = mSiteCalls.rbegin() ; incidentIt != mSiteCalls.rend() ; ++incidentIt )
          {
            incident = incidentIt->second ;
            if( matchCallToSiteIncident(cssCall, incident, false, false, true) )
            {
              break;
            }
          }
        }
      }

      if( incident != NULL )
      {
        if( incidentUse != UseIncidentToHostNewCall )
        {
          if( incidentUse == UseIncidentForCopyingAddress )
          {
            CAli *address = CreateAli()->init();
            if( incident->addressSpecified(AddressType_Ili) ) // if ILI is specified independently, use it directly
            {
              incident->address(address, AddressType_Ili) ;
              INCIDENT(cssCall)->setAddress(ModuleName, mCurrentAgent, address, AddressType_Ili, AddressOrigin_CopyForOutgoing);
            }
          }
          }
          }
    }
}
